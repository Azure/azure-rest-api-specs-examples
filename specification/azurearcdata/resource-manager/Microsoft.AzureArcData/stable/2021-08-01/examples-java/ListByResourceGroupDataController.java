import com.azure.core.util.Context;

/** Samples for DataControllers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/ListByResourceGroupDataController.json
     */
    /**
     * Sample code: Gets all dataControllers in a resource group.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void getsAllDataControllersInAResourceGroup(
        com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager.dataControllers().listByResourceGroup("testrg", Context.NONE);
    }
}
