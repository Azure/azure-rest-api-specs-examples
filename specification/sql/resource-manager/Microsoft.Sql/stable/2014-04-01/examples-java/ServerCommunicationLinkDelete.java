import com.azure.core.util.Context;

/** Samples for ServerCommunicationLinks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServerCommunicationLinkDelete.json
     */
    /**
     * Sample code: Delete a server communication link.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAServerCommunicationLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServerCommunicationLinks()
            .deleteWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645", "link1", Context.NONE);
    }
}
