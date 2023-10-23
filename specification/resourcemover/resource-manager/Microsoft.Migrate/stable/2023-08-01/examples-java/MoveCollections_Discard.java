import com.azure.resourcemanager.resourcemover.models.DiscardRequest;
import java.util.Arrays;

/** Samples for MoveCollections Discard. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Discard.json
     */
    /**
     * Sample code: MoveCollections_Discard.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsDiscard(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveCollections()
            .discard(
                "rg1",
                "movecollection1",
                new DiscardRequest()
                    .withValidateOnly(false)
                    .withMoveResources(
                        Arrays
                            .asList(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")),
                com.azure.core.util.Context.NONE);
    }
}
