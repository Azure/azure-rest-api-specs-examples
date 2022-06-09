```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ProductPolicy CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateProductPolicy.json
     */
    /**
     * Sample code: ApiManagementCreateProductPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateProductPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productPolicies()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                "5702e97e5157a50f48dce801",
                PolicyIdName.POLICY,
                new PolicyContractInner()
                    .withValue(
                        "<policies>\r\n"
                            + "  <inbound>\r\n"
                            + "    <rate-limit calls=\"{{call-count}}\" renewal-period=\"15\"></rate-limit>\r\n"
                            + "    <log-to-eventhub logger-id=\"16\">\r\n"
                            + "                      @( string.Join(\",\", DateTime.UtcNow,"
                            + " context.Deployment.ServiceName, context.RequestId, context.Request.IpAddress,"
                            + " context.Operation.Name) ) \r\n"
                            + "                  </log-to-eventhub>\r\n"
                            + "    <quota-by-key calls=\"40\" counter-key=\"cc\" renewal-period=\"3600\""
                            + " increment-count=\"@(context.Request.Method == &quot;POST&quot; ? 1:2)\" />\r\n"
                            + "    <base />\r\n"
                            + "  </inbound>\r\n"
                            + "  <backend>\r\n"
                            + "    <base />\r\n"
                            + "  </backend>\r\n"
                            + "  <outbound>\r\n"
                            + "    <base />\r\n"
                            + "  </outbound>\r\n"
                            + "</policies>")
                    .withFormat(PolicyContentFormat.XML),
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
