
/**
 * Samples for AzureLargeInstance Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstance_PatchTags.json
     */
    /**
     * Sample code: AzureLargeInstance_Update_Tag.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void
        azureLargeInstanceUpdateTag(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.azureLargeInstances().updateWithResponse("myResourceGroup", "myALInstance", null,
            com.azure.core.util.Context.NONE);
    }
}
