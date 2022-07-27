import com.azure.core.util.Context;

/** Samples for ServerCommunicationLinks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerCommunicationLinkGet.json
     */
    /**
     * Sample code: Get a server communication link.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAServerCommunicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerCommunicationLinks()
            .getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645", "link1", Context.NONE);
    }
}
