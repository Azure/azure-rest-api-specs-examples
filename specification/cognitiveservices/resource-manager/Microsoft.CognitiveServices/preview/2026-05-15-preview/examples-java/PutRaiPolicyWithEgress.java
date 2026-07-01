
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressDefaultAction;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressHeaderOperation;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressHeaderTransform;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressHeaderValueRef;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressManagedIdentityRef;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressMode;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressPolicyConfig;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRewriteTarget;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRule;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleAction;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleActionType;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleMatch;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressRuleType;
import com.azure.resourcemanager.cognitiveservices.models.RaiEgressScheme;
import com.azure.resourcemanager.cognitiveservices.models.RaiPolicyProperties;
import java.util.Arrays;

/**
 * Samples for RaiPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/PutRaiPolicyWithEgress.json
     */
    /**
     * Sample code: PutRaiPolicyWithEgress.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        putRaiPolicyWithEgress(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().define("egress-baseline").withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new RaiPolicyProperties().withBasePolicyName("Microsoft.Default")
                .withEgressPolicy(new RaiEgressPolicyConfig().withMode(RaiEgressMode.ENFORCED)
                    .withDefaultAction(RaiEgressDefaultAction.DENY)
                    .withDescription("Corporate baseline egress policy for sandboxed agents")
                    .withRules(Arrays
                        .asList(
                            new RaiEgressRule().withName("allow-openai").withDescription("Allow traffic to OpenAI API")
                                .withRuleType(RaiEgressRuleType.FQDN)
                                .withMatch(new RaiEgressRuleMatch().withHost("*.openai.com"))
                                .withAction(new RaiEgressRuleAction().withActionType(RaiEgressRuleActionType.ALLOW)),
                            new RaiEgressRule().withName("inject-auth-for-internal")
                                .withDescription("Inject managed identity token for internal services")
                                .withRuleType(RaiEgressRuleType.FQDN)
                                .withMatch(new RaiEgressRuleMatch().withHost("*.internal.contoso.com"))
                                .withAction(new RaiEgressRuleAction().withActionType(RaiEgressRuleActionType.TRANSFORM)
                                    .withHeaders(Arrays.asList(new RaiEgressHeaderTransform()
                                        .withOperation(RaiEgressHeaderOperation.SET).withName("Authorization")
                                        .withValueRef(new RaiEgressHeaderValueRef()
                                            .withManagedIdentityRef(new RaiEgressManagedIdentityRef()
                                                .withResource("https://internal.contoso.com/.default")
                                                .withFormat("Bearer {value}")))))),
                            new RaiEgressRule().withName("rewrite-legacy-api")
                                .withDescription("Rewrite legacy API hostname to new internal endpoint")
                                .withRuleType(RaiEgressRuleType.FQDN)
                                .withMatch(
                                    new RaiEgressRuleMatch().withHost("legacy-api.contoso.com").withPath("/v1/*"))
                                .withAction(new RaiEgressRuleAction().withActionType(RaiEgressRuleActionType.REWRITE)
                                    .withHeaders(Arrays.asList(
                                        new RaiEgressHeaderTransform().withOperation(RaiEgressHeaderOperation.SET)
                                            .withName("X-Forwarded-Host").withValue("legacy-api.contoso.com")))
                                    .withRewrite(new RaiEgressRewriteTarget().withScheme(RaiEgressScheme.HTTPS)
                                        .withHost("api-v2.internal.contoso.com").withPath("/v2/")))))))
            .create();
    }
}
