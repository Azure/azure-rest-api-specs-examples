
/**
 * Samples for JobCredentials Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetJobCredential.json
     */
    /**
     * Sample code: Get a credential.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getACredential(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobCredentials().getWithResponse("group1", "server1", "agent1", "cred1",
            com.azure.core.util.Context.NONE);
    }
}
