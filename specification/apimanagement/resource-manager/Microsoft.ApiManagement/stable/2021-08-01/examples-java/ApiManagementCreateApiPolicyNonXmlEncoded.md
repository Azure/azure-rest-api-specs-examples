Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ApiPolicy CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiPolicyNonXmlEncoded.json
     */
    /**
     * Sample code: ApiManagementCreateApiPolicyNonXmlEncoded.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiPolicyNonXmlEncoded(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiPolicies()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                "5600b57e7e8880006a040001",
                PolicyIdName.POLICY,
                new PolicyContractInner()
                    .withValue(
                        "<policies>\r\n"
                            + "     <inbound>\r\n"
                            + "     <base />\r\n"
                            + "  <set-header name=\"newvalue\" exists-action=\"override\">\r\n"
                            + "   <value>\"@(context.Request.Headers.FirstOrDefault(h => h.Ke==\"Via\"))\" </value>\r\n"
                            + "    </set-header>\r\n"
                            + "  </inbound>\r\n"
                            + "      </policies>")
                    .withFormat(PolicyContentFormat.RAWXML),
                "*",
                Context.NONE);
    }
}
```
