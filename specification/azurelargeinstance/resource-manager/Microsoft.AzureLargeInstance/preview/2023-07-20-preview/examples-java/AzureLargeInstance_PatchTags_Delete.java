
/**
 * Samples for AzureLargeInstance Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_PatchTags_Delete.json
     */
    /**
     * Sample code: AzureLargeInstance_Delete_Tag.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeInstanceDeleteTag(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().updateWithResponse("myResourceGroup", "myALInstance", null,
            com.azure.core.util.Context.NONE);
    }
}
