/** Samples for Appliances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/stable/2022-10-27/examples/AppliancesListByResourceGroup.json
     */
    /**
     * Sample code: List Appliances by resource group.
     *
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void listAppliancesByResourceGroup(
        com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        manager.appliances().listByResourceGroup("testresourcegroup", com.azure.core.util.Context.NONE);
    }
}
