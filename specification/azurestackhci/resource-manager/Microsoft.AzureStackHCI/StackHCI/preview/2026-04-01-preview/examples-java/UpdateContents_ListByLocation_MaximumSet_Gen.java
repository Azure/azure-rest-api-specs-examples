
/**
 * Samples for UpdateContents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/UpdateContents_ListByLocation_MaximumSet_Gen.json
     */
    /**
     * Sample code: UpdateContents_ListByLocation_MaximumSet.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        updateContentsListByLocationMaximumSet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updateContents().list("westus2", com.azure.core.util.Context.NONE);
    }
}
