// Copyright 2020 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stream

import (
	"context"

	"github.com/searKing/golang/go/util"
	"github.com/searKing/golang/go/util/function/binary"
	"github.com/searKing/golang/go/util/function/consumer"
	"github.com/searKing/golang/go/util/function/predicate"
	"github.com/searKing/golang/go/util/optional"
)

/**
 * A sequence of elements supporting sequential and parallel aggregate
 * operations.  The following example illustrates an aggregate operation using
 * {@link Stream} and {@link IntStream}:
 *
 * <pre>{@code
 *     int sum = widgets.stream()
 *                      .filter(w -> w.getColor() == RED)
 *                      .mapToInt(w -> w.getWeight())
 *                      .sum();
 * }</pre>
 *
 * In this example, {@code widgets} is a {@code Collection<Widget>}.  We create
 * a stream of {@code Widget} objects via {@link Collection#stream Collection.stream()},
 * filter it to produce a stream containing only the red widgets, and then
 * transform it into a stream of {@code int} values representing the weight of
 * each red widget. Then this stream is summed to produce a total weight.
 *
 * <p>In addition to {@code Stream}, which is a stream of object references,
 * there are primitive specializations for {@link IntStream}, {@link LongStream},
 * and {@link DoubleStream}, all of which are referred to as "streams" and
 * conform to the characteristics and restrictions described here.
 *
 * <p>To perform a computation, stream
 * <a href="package-summary.html#StreamOps">operations</a> are composed into a
 * <em>stream pipeline</em>.  A stream pipeline consists of a source (which
 * might be an array, a collection, a generator function, an I/O channel,
 * etc), zero or more <em>intermediate operations</em> (which transform a
 * stream into another stream, such as {@link Stream#filter(Predicate)}), and a
 * <em>terminal operation</em> (which produces a result or side-effect, such
 * as {@link Stream#count()} or {@link Stream#forEach(Consumer)}).
 * Streams are lazy; computation on the source data is only performed when the
 * terminal operation is initiated, and source elements are consumed only
 * as needed.
 *
 * <p>A stream implementation is permitted significant latitude in optimizing
 * the computation of the result.  For example, a stream implementation is free
 * to elide operations (or entire stages) from a stream pipeline -- and
 * therefore elide invocation of behavioral parameters -- if it can prove that
 * it would not affect the result of the computation.  This means that
 * side-effects of behavioral parameters may not always be executed and should
 * not be relied upon, unless otherwise specified (such as by the terminal
 * operations {@code forEach} and {@code forEachOrdered}). (For a specific
 * example of such an optimization, see the API note documented on the
 * {@link #count} operation.  For more detail, see the
 * <a href="package-summary.html#SideEffects">side-effects</a> section of the
 * stream package documentation.)
 *
 * <p>Collections and streams, while bearing some superficial similarities,
 * have different goals.  Collections are primarily concerned with the efficient
 * management of, and access to, their elements.  By contrast, streams do not
 * provide a means to directly access or manipulate their elements, and are
 * instead concerned with declaratively describing their source and the
 * computational operations which will be performed in aggregate on that source.
 * However, if the provided stream operations do not offer the desired
 * functionality, the {@link #iterator()} and {@link #spliterator()} operations
 * can be used to perform a controlled traversal.
 *
 * <p>A stream pipeline, like the "widgets" example above, can be viewed as
 * a <em>query</em> on the stream source.  Unless the source was explicitly
 * designed for concurrent modification (such as a {@link ConcurrentHashMap}),
 * unpredictable or erroneous behavior may result from modifying the stream
 * source while it is being queried.
 *
 * <p>Most stream operations accept parameters that describe user-specified
 * behavior, such as the lambda expression {@code w -> w.getWeight()} passed to
 * {@code mapToInt} in the example above.  To preserve correct behavior,
 * these <em>behavioral parameters</em>:
 * <ul>
 * <li>must be <a href="package-summary.html#NonInterference">non-interfering</a>
 * (they do not modify the stream source); and</li>
 * <li>in most cases must be <a href="package-summary.html#Statelessness">stateless</a>
 * (their result should not depend on any state that might change during execution
 * of the stream pipeline).</li>
 * </ul>
 *
 * <p>Such parameters are always instances of a
 * <a href="../function/package-summary.html">functional interface</a> such
 * as {@link java.util.function.Function}, and are often lambda expressions or
 * method references.  Unless otherwise specified these parameters must be
 * <em>non-null</em>.
 *
 * <p>A stream should be operated on (invoking an intermediate or terminal stream
 * operation) only once.  This rules out, for example, "forked" streams, where
 * the same source feeds two or more pipelines, or multiple traversals of the
 * same stream.  A stream implementation may throw {@link IllegalStateException}
 * if it detects that the stream is being reused. However, since some stream
 * operations may return their receiver rather than a new stream object, it may
 * not be possible to detect reuse in all cases.
 *
 * <p>Streams have a {@link #close()} method and implement {@link AutoCloseable}.
 * Operating on a stream after it has been closed will throw {@link IllegalStateException}.
 * Most stream instances do not actually need to be closed after use, as they
 * are backed by collections, arrays, or generating functions, which require no
 * special resource management. Generally, only streams whose source is an IO channel,
 * such as those returned by {@link Files#lines(Path)}, will require closing. If a
 * stream does require closing, it must be opened as a resource within a try-with-resources
 * statement or similar control structure to ensure that it is closed promptly after its
 * operations have completed.
 *
 * <p>Stream pipelines may execute either sequentially or in
 * <a href="package-summary.html#Parallelism">parallel</a>.  This
 * execution mode is a property of the stream.  Streams are created
 * with an initial choice of sequential or parallel execution.  (For example,
 * {@link Collection#stream() Collection.stream()} creates a sequential stream,
 * and {@link Collection#parallelStream() Collection.parallelStream()} creates
 * a parallel one.)  This choice of execution mode may be modified by the
 * {@link #sequential()} or {@link #parallel()} methods, and may be queried with
 * the {@link #isParallel()} method.
 *
 * @param <T> the type of the stream elements
 * @since 1.8
 * @see IntStream
 * @see LongStream
 * @see DoubleStream
 * @see <a href="package-summary.html">java.util.stream</a>
 */
type Stream interface {
	BaseStream
	Filter(ctx context.Context, predicate predicate.Predicater) Stream
	Map(ctx context.Context, f func(interface{}) interface{}) Stream
	Distinct(ctx context.Context, comparator util.Comparator) Stream
	Sorted(lesser func(interface{}, interface{}) int) Stream
	Peek(ctx context.Context, action consumer.Consumer) Stream
	Limit(maxSize int) Stream
	Skip(n int) Stream
	TakeWhile(ctx context.Context, predicate predicate.Predicater) Stream
	TakeUntil(ctx context.Context, predicate predicate.Predicater) Stream
	DropWhile(ctx context.Context, predicate predicate.Predicater) Stream
	DropUntil(ctx context.Context, predicate predicate.Predicater) Stream
	ForEach(ctx context.Context, action consumer.Consumer)
	ForEachOrdered(ctx context.Context, action consumer.Consumer)
	ToSlice() interface{}
	Reduce(ctx context.Context, accumulator binary.BiFunction, identity ...interface{}) optional.Optional
	Min(ctx context.Context, comparator util.Comparator) optional.Optional
	Max(ctx context.Context, comparator util.Comparator) optional.Optional
	Count() int
	AnyMatch(ctx context.Context, predicate predicate.Predicater) bool
	AllMatch(ctx context.Context, predicate predicate.Predicater) bool
	NoneMatch(ctx context.Context, predicate predicate.Predicater) bool
	FindFirst(ctx context.Context, predicate predicate.Predicater) optional.Optional
	FindFirstIndex(ctx context.Context, predicate predicate.Predicater) int
	FindAny(ctx context.Context, predicate predicate.Predicater) optional.Optional
	FindAnyIndex(ctx context.Context, predicate predicate.Predicater) int
	Empty() interface{}
	Of() Stream
	Concat(s2 Stream) Stream
	ConcatWithValue(v interface{}) Stream

	//grammar surger for count
	Size() int
	Length() int
}
