/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/newrelic/resource-manager/NewRelic.Observability/stable/2022-07-01/examples/Accounts_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Accounts_List_MaximumSet_Gen.
     *
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void accountsListMaximumSetGen(
        com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.accounts().list("ruxvg@xqkmdhrnoo.hlmbpm", "egh", com.azure.core.util.Context.NONE);
    }
}
