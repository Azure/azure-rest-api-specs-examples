
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_FromAnElasticSanVolumeSnapshot.json
     */
    /**
     * Sample code: create a managed disk from elastic san volume snapshot.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAManagedDiskFromElasticSanVolumeSnapshot(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US").withCreationData(
                new CreationData().withCreateOption(DiskCreateOption.COPY_FROM_SAN_SNAPSHOT).withElasticSanResourceId(
                    "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ElasticSan/elasticSans/myElasticSan/volumegroups/myElasticSanVolumeGroup/snapshots/myElasticSanVolumeSnapshot")),
            com.azure.core.util.Context.NONE);
    }
}
