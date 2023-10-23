import com.azure.resourcemanager.resourcemover.models.Identity;
import com.azure.resourcemanager.resourcemover.models.MoveCollectionProperties;
import com.azure.resourcemanager.resourcemover.models.MoveType;
import com.azure.resourcemanager.resourcemover.models.ResourceIdentityType;

/** Samples for MoveCollections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Create.json
     */
    /**
     * Sample code: MoveCollections_Create.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsCreate(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveCollections()
            .define("movecollection1")
            .withRegion("eastus2")
            .withExistingResourceGroup("rg1")
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(
                new MoveCollectionProperties()
                    .withSourceRegion("eastus")
                    .withTargetRegion("westus")
                    .withMoveType(MoveType.REGION_TO_REGION))
            .create();
    }
}
