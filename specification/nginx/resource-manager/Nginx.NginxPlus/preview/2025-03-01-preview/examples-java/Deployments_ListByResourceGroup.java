
/**
 * Samples for Deployments ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Deployments_ListByResourceGroup.json
     */
    /**
     * Sample code: Deployments_ListByResourceGroup.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void deploymentsListByResourceGroup(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.deployments().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
