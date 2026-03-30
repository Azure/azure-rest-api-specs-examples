
/**
 * Samples for ResourceHealthMetadata ListBySite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListResourceHealthMetadataBySite.json
     */
    /**
     * Sample code: List ResourceHealthMetadata for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listResourceHealthMetadataForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getResourceHealthMetadatas().listBySite("Default-Web-NorthCentralUS",
            "newsiteinnewASE-NCUS", com.azure.core.util.Context.NONE);
    }
}
