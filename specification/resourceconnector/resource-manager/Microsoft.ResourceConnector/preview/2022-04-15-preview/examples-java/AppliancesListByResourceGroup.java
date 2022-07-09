import com.azure.core.util.Context;

/** Samples for Appliances ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourceconnector/resource-manager/Microsoft.ResourceConnector/preview/2022-04-15-preview/examples/AppliancesListByResourceGroup.json
     */
    /**
     * Sample code: List Appliances by resource group.
     *
     * @param manager Entry point to AppliancesManager.
     */
    public static void listAppliancesByResourceGroup(
        com.azure.resourcemanager.resourceconnector.AppliancesManager manager) {
        manager.appliances().listByResourceGroup("testresourcegroup", Context.NONE);
    }
}
