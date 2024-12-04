
/**
 * Samples for AzureLargeStorageInstance Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeStorageInstance_PatchTags_Delete.json
     */
    /**
     * Sample code: AzureLargeStorageInstance_Delete_Tag.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeStorageInstanceDeleteTag(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeStorageInstances().updateWithResponse("myResourceGroup", "myALSInstance", null,
            com.azure.core.util.Context.NONE);
    }
}
