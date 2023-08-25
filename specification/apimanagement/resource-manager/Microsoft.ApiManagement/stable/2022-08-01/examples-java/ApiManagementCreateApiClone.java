import com.azure.resourcemanager.apimanagement.models.Protocol;
import java.util.Arrays;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiClone.json
     */
    /**
     * Sample code: ApiManagementCreateApiClone.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiClone(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("echo-api2")
            .withExistingService("rg1", "apimService1")
            .withSourceApiId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/58a4aeac497000007d040001")
            .withDisplayName("Echo API2")
            .withServiceUrl("http://echoapi.cloudapp.net/api")
            .withPath("echo2")
            .withProtocols(Arrays.asList(Protocol.HTTP, Protocol.HTTPS))
            .withDescription("Copy of Existing Echo Api including Operations.")
            .withIsCurrent(true)
            .withSubscriptionRequired(true)
            .create();
    }
}
