
/**
 * Samples for AzureLargeStorageInstance GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeStorageInstance_Get.json
     */
    /**
     * Sample code: AzureLargeStorageInstance_Get.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeStorageInstanceGet(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeStorageInstances().getByResourceGroupWithResponse("myResourceGroup",
            "myAzureLargeStorageInstance", com.azure.core.util.Context.NONE);
    }
}
