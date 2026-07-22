
import com.azure.resourcemanager.compute.fluent.models.SnapshotInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Snapshots CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-02/snapshotExamples/Snapshot_Create_ByImportingAnUnmanagedBlobFromTheSameSubscription.json
     */
    /**
     * Sample code: Create a snapshot by importing an unmanaged blob from the same subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createASnapshotByImportingAnUnmanagedBlobFromTheSameSubscription(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSnapshots().createOrUpdate("myResourceGroup", "mySnapshot1",
            new SnapshotInner().withLocation("West US")
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.IMPORT)
                    .withSourceUri("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd")),
            com.azure.core.util.Context.NONE);
    }
}
