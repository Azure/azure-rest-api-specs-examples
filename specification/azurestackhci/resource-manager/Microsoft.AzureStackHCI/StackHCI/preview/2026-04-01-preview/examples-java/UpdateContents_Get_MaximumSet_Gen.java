
/**
 * Samples for UpdateContents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/UpdateContents_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: UpdateContents_Get_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        updateContentsGetMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateContents().getWithResponse("westus2", "12.2510.0.1", com.azure.core.util.Context.NONE);
    }
}
