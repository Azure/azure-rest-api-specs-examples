
/**
 * Samples for ResourceHealthMetadata GetBySite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetResourceHealthMetadataBySite.json
     */
    /**
     * Sample code: Get ResourceHealthMetadata.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getResourceHealthMetadata(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceHealthMetadatas().getBySiteWithResponse("Default-Web-NorthCentralUS",
            "newsiteinnewASE-NCUS", com.azure.core.util.Context.NONE);
    }
}
