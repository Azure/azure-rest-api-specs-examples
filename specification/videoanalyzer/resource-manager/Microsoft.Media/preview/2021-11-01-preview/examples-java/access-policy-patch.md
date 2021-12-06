Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.4/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.AccessPolicyEccAlgo;
import com.azure.resourcemanager.videoanalyzer.models.AccessPolicyEntity;
import com.azure.resourcemanager.videoanalyzer.models.AccessPolicyRsaAlgo;
import com.azure.resourcemanager.videoanalyzer.models.EccTokenKey;
import com.azure.resourcemanager.videoanalyzer.models.JwtAuthentication;
import com.azure.resourcemanager.videoanalyzer.models.RsaTokenKey;
import com.azure.resourcemanager.videoanalyzer.models.TokenClaim;
import java.util.Arrays;

/** Samples for AccessPolicies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/access-policy-patch.json
     */
    /**
     * Sample code: Update access policy entity.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void updateAccessPolicyEntity(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        AccessPolicyEntity resource =
            manager
                .accessPolicies()
                .getWithResponse("testrg", "testaccount2", "accessPolicyName1", Context.NONE)
                .getValue();
        resource
            .update()
            .withAuthentication(
                new JwtAuthentication()
                    .withIssuers(Arrays.asList("issuer1", "issuer2"))
                    .withAudiences(Arrays.asList("audience1"))
                    .withClaims(
                        Arrays
                            .asList(
                                new TokenClaim().withName("claimname1").withValue("claimvalue1"),
                                new TokenClaim().withName("claimname2").withValue("claimvalue2")))
                    .withKeys(
                        Arrays
                            .asList(
                                new RsaTokenKey()
                                    .withKid("123")
                                    .withAlg(AccessPolicyRsaAlgo.RS256)
                                    .withN("YmFzZTY0IQ==")
                                    .withE("ZLFzZTY0IQ=="),
                                new EccTokenKey()
                                    .withKid("124")
                                    .withAlg(AccessPolicyEccAlgo.ES256)
                                    .withX("XX==")
                                    .withY("YY=="))))
            .apply();
    }
}
```
