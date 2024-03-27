
/**
 * Samples for Deployments ListByDeviceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetDeployments.json
     */
    /**
     * Sample code: Deployments_ListByDeviceGroup.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deploymentsListByDeviceGroup(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.deployments().listByDeviceGroup("MyResourceGroup1", "MyCatalog1", "MyProduct1", "myDeviceGroup1", null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
