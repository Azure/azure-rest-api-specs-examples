
/**
 * Samples for JobCredentials Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetJobCredential.json
     */
    /**
     * Sample code: Get a credential.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getACredential(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobCredentials().getWithResponse("group1", "server1", "agent1",
            "cred1", com.azure.core.util.Context.NONE);
    }
}
