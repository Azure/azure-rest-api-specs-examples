import com.azure.core.util.Context;

/** Samples for Linker Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/Link.json
     */
    /**
     * Sample code: Link.
     *
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void link(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager
            .linkers()
            .getWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
                "linkName",
                Context.NONE);
    }
}
