/** Samples for Deployments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/DeleteDeployment.json
     */
    /**
     * Sample code: Deployments_Delete.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deploymentsDelete(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .deployments()
            .delete(
                "MyResourceGroup1",
                "MyCatalog1",
                "MyProductName1",
                "DeviceGroupName1",
                "MyDeploymentName1",
                com.azure.core.util.Context.NONE);
    }
}
