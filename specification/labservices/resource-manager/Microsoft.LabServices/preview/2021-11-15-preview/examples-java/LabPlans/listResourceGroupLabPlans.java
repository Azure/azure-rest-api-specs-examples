import com.azure.core.util.Context;

/** Samples for LabPlans ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/listResourceGroupLabPlans.json
     */
    /**
     * Sample code: listResourceGroupLabPlans.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listResourceGroupLabPlans(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labPlans().listByResourceGroup("testrg123", Context.NONE);
    }
}
