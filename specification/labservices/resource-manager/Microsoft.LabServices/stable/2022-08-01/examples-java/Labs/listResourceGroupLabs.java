import com.azure.core.util.Context;

/** Samples for Labs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Labs/listResourceGroupLabs.json
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
