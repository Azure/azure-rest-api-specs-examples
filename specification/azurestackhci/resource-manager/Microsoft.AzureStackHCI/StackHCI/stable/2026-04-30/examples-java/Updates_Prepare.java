
/**
 * Samples for Updates Prepare.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/Updates_Prepare.json
     */
    /**
     * Sample code: Prepare Update.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void prepareUpdate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.updates().prepare("testrg", "testcluster", "Microsoft4.2203.2.32", com.azure.core.util.Context.NONE);
    }
}
