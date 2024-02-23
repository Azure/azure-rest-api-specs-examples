
/**
 * Samples for Apis List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/Apis_List.json
     */
    /**
     * Sample code: Apis_ListByWorkspace.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void apisListByWorkspace(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.apis().list("contoso-resources", "contoso", "default", null, com.azure.core.util.Context.NONE);
    }
}
