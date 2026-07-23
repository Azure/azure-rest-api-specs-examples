
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.ExtendedLocation;
import com.azure.resourcemanager.compute.models.ExtendedLocationTypes;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Create_InExtendedLocation.json
     */
    /**
     * Sample code: create an empty managed disk in extended location.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAnEmptyManagedDiskInExtendedLocation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US")
                .withExtendedLocation(
                    new ExtendedLocation().withName("{edge-zone-id}").withType(ExtendedLocationTypes.EDGE_ZONE))
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.EMPTY)).withDiskSizeGB(200),
            com.azure.core.util.Context.NONE);
    }
}
