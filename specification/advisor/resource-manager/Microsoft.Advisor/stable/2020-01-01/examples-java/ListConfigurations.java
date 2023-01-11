/** Samples for Configurations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListConfigurations.json
     */
    /**
     * Sample code: GetConfigurations.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void getConfigurations(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager.configurations().listByResourceGroup("resourceGroup", com.azure.core.util.Context.NONE);
    }
}
