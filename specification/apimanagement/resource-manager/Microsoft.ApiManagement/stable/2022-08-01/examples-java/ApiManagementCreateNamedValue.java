import java.util.Arrays;

/** Samples for NamedValue CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateNamedValue.json
     */
    /**
     * Sample code: ApiManagementCreateNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .namedValues()
            .define("testprop2")
            .withExistingService("rg1", "apimService1")
            .withTags(Arrays.asList("foo", "bar"))
            .withDisplayName("prop3name")
            .withValue("propValue")
            .withSecret(false)
            .create();
    }
}
