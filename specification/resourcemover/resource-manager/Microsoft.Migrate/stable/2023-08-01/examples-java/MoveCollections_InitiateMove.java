import com.azure.resourcemanager.resourcemover.models.ResourceMoveRequest;
import java.util.Arrays;

/** Samples for MoveCollections InitiateMove. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_InitiateMove.json
     */
    /**
     * Sample code: MoveCollections_InitiateMove.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsInitiateMove(
        com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveCollections()
            .initiateMove(
                "rg1",
                "movecollection1",
                new ResourceMoveRequest()
                    .withValidateOnly(false)
                    .withMoveResources(
                        Arrays
                            .asList(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")),
                com.azure.core.util.Context.NONE);
    }
}
