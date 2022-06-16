import com.azure.core.util.Context;

/** Samples for Labs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Labs/listResourceGroupLabs.json
     */
    /**
     * Sample code: listResourceGroupLabs.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void listResourceGroupLabs(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.labs().listByResourceGroup("testrg123", Context.NONE);
    }
}
