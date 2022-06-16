import com.azure.core.util.Context;

/** Samples for EdgeNodes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/EdgeNodes_List.json
     */
    /**
     * Sample code: EdgeNodes_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void edgeNodesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getEdgeNodes().list(Context.NONE);
    }
}
