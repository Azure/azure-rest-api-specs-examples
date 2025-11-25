
/**
 * Samples for Accounts List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/
     * Accounts_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Accounts_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void accountsListMinimumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.accounts().list("ruxvg@xqkmdhrnoo.hlmbpm", "egh", com.azure.core.util.Context.NONE);
    }
}
