/** Samples for Clusters Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/CreateCluster.json
     */
    /**
     * Sample code: Create cluster.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createCluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .clusters()
            .define("myCluster")
            .withRegion("East US")
            .withExistingResourceGroup("test-rg")
            .withCloudManagementEndpoint("https://98294836-31be-4668-aeae-698667faf99b.waconazure.com")
            .withAadClientId("24a6e53d-04e5-44d2-b7cc-1b732a847dfc")
            .withAadTenantId("7e589cc1-a8b6-4dff-91bd-5ec0fa18db94")
            .create();
    }
}
