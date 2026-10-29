
/**
 * Samples for EdgeNodes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/EdgeNodes_List.json
     */
    /**
     * Sample code: EdgeNodes_List.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void edgeNodesList(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEdgeNodes().list(com.azure.core.util.Context.NONE);
    }
}
