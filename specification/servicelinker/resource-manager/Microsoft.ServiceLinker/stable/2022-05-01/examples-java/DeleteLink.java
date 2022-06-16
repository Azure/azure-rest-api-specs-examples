import com.azure.core.util.Context;

/** Samples for Linker Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/DeleteLink.json
     */
    /**
     * Sample code: DeleteLink.
     *
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void deleteLink(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager
            .linkers()
            .delete(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
                "linkName",
                Context.NONE);
    }
}
