/** Samples for DataControllers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/stable/2021-08-01/examples/GetDataController.json
     */
    /**
     * Sample code: Get a data controller.
     *
     * @param manager Entry point to AzureArcDataManager.
     */
    public static void getADataController(com.azure.resourcemanager.azurearcdata.AzureArcDataManager manager) {
        manager
            .dataControllers()
            .getByResourceGroupWithResponse("testrg", "testdataController", com.azure.core.util.Context.NONE);
    }
}
