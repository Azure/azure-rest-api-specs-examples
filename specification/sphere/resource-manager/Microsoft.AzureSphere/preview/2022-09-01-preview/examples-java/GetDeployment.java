/** Samples for Deployments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetDeployment.json
     */
    /**
     * Sample code: Deployments_Get.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deploymentsGet(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .deployments()
            .getWithResponse(
                "MyResourceGroup1",
                "MyCatalog1",
                "MyProduct1",
                "myDeviceGroup1",
                "MyDeployment1",
                com.azure.core.util.Context.NONE);
    }
}
