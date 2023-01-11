/** Samples for Suppressions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListSuppressions.json
     */
    /**
     * Sample code: ListSuppressions.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void listSuppressions(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager.suppressions().list(null, null, com.azure.core.util.Context.NONE);
    }
}
