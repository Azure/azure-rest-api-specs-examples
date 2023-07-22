/** Samples for Deployments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PutDeployment.json
     */
    /**
     * Sample code: Deployments_CreateOrUpdate.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deploymentsCreateOrUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .deployments()
            .define("MyDeployment1")
            .withExistingDeviceGroup("MyResourceGroup1", "MyCatalog1", "MyProduct1", "myDeviceGroup1")
            .create();
    }
}
