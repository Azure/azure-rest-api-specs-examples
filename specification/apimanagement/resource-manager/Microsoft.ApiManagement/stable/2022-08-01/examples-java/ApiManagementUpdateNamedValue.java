import com.azure.resourcemanager.apimanagement.models.NamedValueContract;
import java.util.Arrays;

/** Samples for NamedValue Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateNamedValue.json
     */
    /**
     * Sample code: ApiManagementUpdateNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        NamedValueContract resource =
            manager
                .namedValues()
                .getWithResponse("rg1", "apimService1", "testprop2", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(Arrays.asList("foo", "bar2"))
            .withDisplayName("prop3name")
            .withValue("propValue")
            .withSecret(false)
            .withIfMatch("*")
            .apply();
    }
}
