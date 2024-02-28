
/**
 * Samples for AzureLargeInstance GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_Get.json
     */
    /**
     * Sample code: AzureLargeInstance_Get.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void azureLargeInstanceGet(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().getByResourceGroupWithResponse("myResourceGroup", "myAzureLargeInstance",
            com.azure.core.util.Context.NONE);
    }
}
