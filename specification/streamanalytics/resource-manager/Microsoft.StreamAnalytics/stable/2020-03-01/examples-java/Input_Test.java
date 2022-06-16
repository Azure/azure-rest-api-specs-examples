import com.azure.core.util.Context;

/** Samples for Inputs Test. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Test.json
     */
    /**
     * Sample code: Test the connection for an input.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void testTheConnectionForAnInput(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager.inputs().test("sjrg8440", "sj9597", "input7225", null, Context.NONE);
    }
}
