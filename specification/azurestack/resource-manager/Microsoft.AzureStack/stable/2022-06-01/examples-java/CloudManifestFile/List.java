/** Samples for CloudManifestFile List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/CloudManifestFile/List.json
     */
    /**
     * Sample code: Returns the properties of a cloud specific manifest file with latest version.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsThePropertiesOfACloudSpecificManifestFileWithLatestVersion(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager.cloudManifestFiles().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
