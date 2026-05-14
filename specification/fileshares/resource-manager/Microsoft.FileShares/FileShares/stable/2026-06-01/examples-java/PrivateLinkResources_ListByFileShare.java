
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/PrivateLinkResources_ListByFileShare.json
     */
    /**
     * Sample code: List PrivateLinkResources.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void listPrivateLinkResources(com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.privateLinkResources().list("res4303", "testfileshare01", com.azure.core.util.Context.NONE);
    }
}
