import com.azure.resourcemanager.databoxedge.models.AzureContainerDataFormat;

/** Samples for Containers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ContainerPut.json
     */
    /**
     * Sample code: ContainerPut.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void containerPut(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .containers()
            .define("blobcontainer1")
            .withExistingStorageAccount("testedgedevice", "storageaccount1", "GroupForEdgeAutomation")
            .withDataFormat(AzureContainerDataFormat.BLOCK_BLOB)
            .create();
    }
}
