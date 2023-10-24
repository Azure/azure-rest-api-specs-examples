import com.azure.resourcemanager.resourcemover.models.PrepareRequest;
import java.util.Arrays;

/** Samples for MoveCollections Prepare. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Prepare.json
     */
    /**
     * Sample code: MoveCollections_Prepare.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsPrepare(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveCollections()
            .prepare(
                "rg1",
                "movecollection1",
                new PrepareRequest()
                    .withValidateOnly(false)
                    .withMoveResources(
                        Arrays
                            .asList(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")),
                com.azure.core.util.Context.NONE);
    }
}
