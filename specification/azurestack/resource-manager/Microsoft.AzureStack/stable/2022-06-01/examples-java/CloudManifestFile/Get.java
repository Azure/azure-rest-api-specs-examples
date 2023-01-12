/** Samples for CloudManifestFile Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/CloudManifestFile/Get.json
     */
    /**
     * Sample code: Returns the properties of a cloud specific manifest file.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsThePropertiesOfACloudSpecificManifestFile(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager.cloudManifestFiles().getWithResponse("latest", null, com.azure.core.util.Context.NONE);
    }
}
