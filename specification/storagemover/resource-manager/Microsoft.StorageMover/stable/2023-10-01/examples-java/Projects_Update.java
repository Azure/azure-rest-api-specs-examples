import com.azure.resourcemanager.storagemover.models.Project;

/** Samples for Projects Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/Projects_Update.json
     */
    /**
     * Sample code: Projects_Update.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void projectsUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        Project resource =
            manager
                .projects()
                .getWithResponse(
                    "examples-rg",
                    "examples-storageMoverName",
                    "examples-projectName",
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withDescription("Example Project Description").apply();
    }
}
