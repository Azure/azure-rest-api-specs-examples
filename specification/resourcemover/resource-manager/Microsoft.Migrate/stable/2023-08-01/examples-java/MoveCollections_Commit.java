import com.azure.resourcemanager.resourcemover.models.CommitRequest;
import java.util.Arrays;

/** Samples for MoveCollections Commit. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Commit.json
     */
    /**
     * Sample code: MoveCollections_Commit.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void moveCollectionsCommit(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveCollections()
            .commit(
                "rg1",
                "movecollection1",
                new CommitRequest()
                    .withValidateOnly(false)
                    .withMoveResources(
                        Arrays
                            .asList(
                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")),
                com.azure.core.util.Context.NONE);
    }
}
