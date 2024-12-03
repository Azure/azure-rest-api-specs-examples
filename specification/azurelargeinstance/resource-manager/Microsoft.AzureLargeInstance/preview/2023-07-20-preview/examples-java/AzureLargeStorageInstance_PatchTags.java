
/**
 * Samples for AzureLargeStorageInstance Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeStorageInstance_PatchTags.json
     */
    /**
     * Sample code: AzureLargeStorageInstance_Update_Tag.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeStorageInstanceUpdateTag(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeStorageInstances().updateWithResponse("myResourceGroup", "myALSInstance", null,
            com.azure.core.util.Context.NONE);
    }
}
