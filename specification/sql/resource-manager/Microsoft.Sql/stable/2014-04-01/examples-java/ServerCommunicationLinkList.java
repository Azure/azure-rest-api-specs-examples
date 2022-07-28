import com.azure.core.util.Context;

/** Samples for ServerCommunicationLinks ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerCommunicationLinkList.json
     */
    /**
     * Sample code: List server communication links.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServerCommunicationLinks(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerCommunicationLinks()
            .listByServer("sqlcrudtest-7398", "sqlcrudtest-4645", Context.NONE);
    }
}
