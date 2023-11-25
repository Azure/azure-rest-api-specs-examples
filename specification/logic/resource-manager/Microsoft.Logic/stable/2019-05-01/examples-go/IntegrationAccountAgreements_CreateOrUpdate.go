package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_CreateOrUpdate.json
func ExampleIntegrationAccountAgreementsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountAgreementsClient().CreateOrUpdate(ctx, "testResourceGroup", "testIntegrationAccount", "testAgreement", armlogic.IntegrationAccountAgreement{
		Location: to.Ptr("westus"),
		Tags: map[string]*string{
			"IntegrationAccountAgreement": to.Ptr("<IntegrationAccountAgreementName>"),
		},
		Properties: &armlogic.IntegrationAccountAgreementProperties{
			AgreementType: to.Ptr(armlogic.AgreementTypeAS2),
			Content: &armlogic.AgreementContent{
				AS2: &armlogic.AS2AgreementContent{
					ReceiveAgreement: &armlogic.AS2OneWayAgreement{
						ProtocolSettings: &armlogic.AS2ProtocolSettings{
							AcknowledgementConnectionSettings: &armlogic.AS2AcknowledgementConnectionSettings{
								IgnoreCertificateNameMismatch: to.Ptr(true),
								KeepHTTPConnectionAlive:       to.Ptr(true),
								SupportHTTPStatusCodeContinue: to.Ptr(true),
								UnfoldHTTPHeaders:             to.Ptr(true),
							},
							EnvelopeSettings: &armlogic.AS2EnvelopeSettings{
								AutogenerateFileName:                    to.Ptr(true),
								FileNameTemplate:                        to.Ptr("Test"),
								MessageContentType:                      to.Ptr("text/plain"),
								SuspendMessageOnFileNameGenerationError: to.Ptr(true),
								TransmitFileNameInMimeHeader:            to.Ptr(true),
							},
							ErrorSettings: &armlogic.AS2ErrorSettings{
								ResendIfMDNNotReceived:  to.Ptr(true),
								SuspendDuplicateMessage: to.Ptr(true),
							},
							MdnSettings: &armlogic.AS2MdnSettings{
								DispositionNotificationTo:  to.Ptr("http://tempuri.org"),
								MdnText:                    to.Ptr("Sample"),
								MicHashingAlgorithm:        to.Ptr(armlogic.HashingAlgorithmSHA1),
								NeedMDN:                    to.Ptr(true),
								ReceiptDeliveryURL:         to.Ptr("http://tempuri.org"),
								SendInboundMDNToMessageBox: to.Ptr(true),
								SendMDNAsynchronously:      to.Ptr(true),
								SignMDN:                    to.Ptr(true),
								SignOutboundMDNIfOptional:  to.Ptr(true),
							},
							MessageConnectionSettings: &armlogic.AS2MessageConnectionSettings{
								IgnoreCertificateNameMismatch: to.Ptr(true),
								KeepHTTPConnectionAlive:       to.Ptr(true),
								SupportHTTPStatusCodeContinue: to.Ptr(true),
								UnfoldHTTPHeaders:             to.Ptr(true),
							},
							SecuritySettings: &armlogic.AS2SecuritySettings{
								EnableNRRForInboundDecodedMessages:  to.Ptr(true),
								EnableNRRForInboundEncodedMessages:  to.Ptr(true),
								EnableNRRForInboundMDN:              to.Ptr(true),
								EnableNRRForOutboundDecodedMessages: to.Ptr(true),
								EnableNRRForOutboundEncodedMessages: to.Ptr(true),
								EnableNRRForOutboundMDN:             to.Ptr(true),
								OverrideGroupSigningCertificate:     to.Ptr(false),
							},
							ValidationSettings: &armlogic.AS2ValidationSettings{
								CheckCertificateRevocationListOnReceive: to.Ptr(true),
								CheckCertificateRevocationListOnSend:    to.Ptr(true),
								CheckDuplicateMessage:                   to.Ptr(true),
								CompressMessage:                         to.Ptr(true),
								EncryptMessage:                          to.Ptr(false),
								EncryptionAlgorithm:                     to.Ptr(armlogic.EncryptionAlgorithmAES128),
								InterchangeDuplicatesValidityDays:       to.Ptr[int32](100),
								OverrideMessageProperties:               to.Ptr(true),
								SignMessage:                             to.Ptr(false),
							},
						},
						ReceiverBusinessIdentity: &armlogic.BusinessIdentity{
							Qualifier: to.Ptr("ZZ"),
							Value:     to.Ptr("ZZ"),
						},
						SenderBusinessIdentity: &armlogic.BusinessIdentity{
							Qualifier: to.Ptr("AA"),
							Value:     to.Ptr("AA"),
						},
					},
					SendAgreement: &armlogic.AS2OneWayAgreement{
						ProtocolSettings: &armlogic.AS2ProtocolSettings{
							AcknowledgementConnectionSettings: &armlogic.AS2AcknowledgementConnectionSettings{
								IgnoreCertificateNameMismatch: to.Ptr(true),
								KeepHTTPConnectionAlive:       to.Ptr(true),
								SupportHTTPStatusCodeContinue: to.Ptr(true),
								UnfoldHTTPHeaders:             to.Ptr(true),
							},
							EnvelopeSettings: &armlogic.AS2EnvelopeSettings{
								AutogenerateFileName:                    to.Ptr(true),
								FileNameTemplate:                        to.Ptr("Test"),
								MessageContentType:                      to.Ptr("text/plain"),
								SuspendMessageOnFileNameGenerationError: to.Ptr(true),
								TransmitFileNameInMimeHeader:            to.Ptr(true),
							},
							ErrorSettings: &armlogic.AS2ErrorSettings{
								ResendIfMDNNotReceived:  to.Ptr(true),
								SuspendDuplicateMessage: to.Ptr(true),
							},
							MdnSettings: &armlogic.AS2MdnSettings{
								DispositionNotificationTo:  to.Ptr("http://tempuri.org"),
								MdnText:                    to.Ptr("Sample"),
								MicHashingAlgorithm:        to.Ptr(armlogic.HashingAlgorithmSHA1),
								NeedMDN:                    to.Ptr(true),
								ReceiptDeliveryURL:         to.Ptr("http://tempuri.org"),
								SendInboundMDNToMessageBox: to.Ptr(true),
								SendMDNAsynchronously:      to.Ptr(true),
								SignMDN:                    to.Ptr(true),
								SignOutboundMDNIfOptional:  to.Ptr(true),
							},
							MessageConnectionSettings: &armlogic.AS2MessageConnectionSettings{
								IgnoreCertificateNameMismatch: to.Ptr(true),
								KeepHTTPConnectionAlive:       to.Ptr(true),
								SupportHTTPStatusCodeContinue: to.Ptr(true),
								UnfoldHTTPHeaders:             to.Ptr(true),
							},
							SecuritySettings: &armlogic.AS2SecuritySettings{
								EnableNRRForInboundDecodedMessages:  to.Ptr(true),
								EnableNRRForInboundEncodedMessages:  to.Ptr(true),
								EnableNRRForInboundMDN:              to.Ptr(true),
								EnableNRRForOutboundDecodedMessages: to.Ptr(true),
								EnableNRRForOutboundEncodedMessages: to.Ptr(true),
								EnableNRRForOutboundMDN:             to.Ptr(true),
								OverrideGroupSigningCertificate:     to.Ptr(false),
							},
							ValidationSettings: &armlogic.AS2ValidationSettings{
								CheckCertificateRevocationListOnReceive: to.Ptr(true),
								CheckCertificateRevocationListOnSend:    to.Ptr(true),
								CheckDuplicateMessage:                   to.Ptr(true),
								CompressMessage:                         to.Ptr(true),
								EncryptMessage:                          to.Ptr(false),
								EncryptionAlgorithm:                     to.Ptr(armlogic.EncryptionAlgorithmAES128),
								InterchangeDuplicatesValidityDays:       to.Ptr[int32](100),
								OverrideMessageProperties:               to.Ptr(true),
								SignMessage:                             to.Ptr(false),
							},
						},
						ReceiverBusinessIdentity: &armlogic.BusinessIdentity{
							Qualifier: to.Ptr("AA"),
							Value:     to.Ptr("AA"),
						},
						SenderBusinessIdentity: &armlogic.BusinessIdentity{
							Qualifier: to.Ptr("ZZ"),
							Value:     to.Ptr("ZZ"),
						},
					},
				},
			},
			GuestIdentity: &armlogic.BusinessIdentity{
				Qualifier: to.Ptr("AA"),
				Value:     to.Ptr("AA"),
			},
			GuestPartner: to.Ptr("GuestPartner"),
			HostIdentity: &armlogic.BusinessIdentity{
				Qualifier: to.Ptr("ZZ"),
				Value:     to.Ptr("ZZ"),
			},
			HostPartner: to.Ptr("HostPartner"),
			Metadata:    map[string]any{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountAgreement = armlogic.IntegrationAccountAgreement{
	// 	Name: to.Ptr("<IntegrationAccountAgreementName>"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/agreements"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationAccounts/IntegrationAccount4533/agreements/<IntegrationAccountAgreementName>"),
	// 	Properties: &armlogic.IntegrationAccountAgreementProperties{
	// 		AgreementType: to.Ptr(armlogic.AgreementTypeAS2),
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T22:32:54.294Z"); return t}()),
	// 		Content: &armlogic.AgreementContent{
	// 			AS2: &armlogic.AS2AgreementContent{
	// 				ReceiveAgreement: &armlogic.AS2OneWayAgreement{
	// 					ProtocolSettings: &armlogic.AS2ProtocolSettings{
	// 						AcknowledgementConnectionSettings: &armlogic.AS2AcknowledgementConnectionSettings{
	// 							IgnoreCertificateNameMismatch: to.Ptr(true),
	// 							KeepHTTPConnectionAlive: to.Ptr(true),
	// 							SupportHTTPStatusCodeContinue: to.Ptr(true),
	// 							UnfoldHTTPHeaders: to.Ptr(true),
	// 						},
	// 						EnvelopeSettings: &armlogic.AS2EnvelopeSettings{
	// 							AutogenerateFileName: to.Ptr(true),
	// 							FileNameTemplate: to.Ptr("Test"),
	// 							MessageContentType: to.Ptr("text/plain"),
	// 							SuspendMessageOnFileNameGenerationError: to.Ptr(true),
	// 							TransmitFileNameInMimeHeader: to.Ptr(true),
	// 						},
	// 						ErrorSettings: &armlogic.AS2ErrorSettings{
	// 							ResendIfMDNNotReceived: to.Ptr(true),
	// 							SuspendDuplicateMessage: to.Ptr(true),
	// 						},
	// 						MdnSettings: &armlogic.AS2MdnSettings{
	// 							DispositionNotificationTo: to.Ptr("http://tempuri.org"),
	// 							MdnText: to.Ptr("Sample"),
	// 							MicHashingAlgorithm: to.Ptr(armlogic.HashingAlgorithmSHA1),
	// 							NeedMDN: to.Ptr(true),
	// 							ReceiptDeliveryURL: to.Ptr("http://tempuri.org"),
	// 							SendInboundMDNToMessageBox: to.Ptr(true),
	// 							SendMDNAsynchronously: to.Ptr(true),
	// 							SignMDN: to.Ptr(true),
	// 							SignOutboundMDNIfOptional: to.Ptr(true),
	// 						},
	// 						MessageConnectionSettings: &armlogic.AS2MessageConnectionSettings{
	// 							IgnoreCertificateNameMismatch: to.Ptr(true),
	// 							KeepHTTPConnectionAlive: to.Ptr(true),
	// 							SupportHTTPStatusCodeContinue: to.Ptr(true),
	// 							UnfoldHTTPHeaders: to.Ptr(true),
	// 						},
	// 						SecuritySettings: &armlogic.AS2SecuritySettings{
	// 							EnableNRRForInboundDecodedMessages: to.Ptr(true),
	// 							EnableNRRForInboundEncodedMessages: to.Ptr(true),
	// 							EnableNRRForInboundMDN: to.Ptr(true),
	// 							EnableNRRForOutboundDecodedMessages: to.Ptr(true),
	// 							EnableNRRForOutboundEncodedMessages: to.Ptr(true),
	// 							EnableNRRForOutboundMDN: to.Ptr(true),
	// 							OverrideGroupSigningCertificate: to.Ptr(false),
	// 						},
	// 						ValidationSettings: &armlogic.AS2ValidationSettings{
	// 							CheckCertificateRevocationListOnReceive: to.Ptr(true),
	// 							CheckCertificateRevocationListOnSend: to.Ptr(true),
	// 							CheckDuplicateMessage: to.Ptr(true),
	// 							CompressMessage: to.Ptr(true),
	// 							EncryptMessage: to.Ptr(false),
	// 							EncryptionAlgorithm: to.Ptr(armlogic.EncryptionAlgorithmAES128),
	// 							InterchangeDuplicatesValidityDays: to.Ptr[int32](100),
	// 							OverrideMessageProperties: to.Ptr(true),
	// 							SignMessage: to.Ptr(false),
	// 						},
	// 					},
	// 					ReceiverBusinessIdentity: &armlogic.BusinessIdentity{
	// 						Qualifier: to.Ptr("ZZ"),
	// 						Value: to.Ptr("ZZ"),
	// 					},
	// 					SenderBusinessIdentity: &armlogic.BusinessIdentity{
	// 						Qualifier: to.Ptr("AA"),
	// 						Value: to.Ptr("AA"),
	// 					},
	// 				},
	// 				SendAgreement: &armlogic.AS2OneWayAgreement{
	// 					ProtocolSettings: &armlogic.AS2ProtocolSettings{
	// 						AcknowledgementConnectionSettings: &armlogic.AS2AcknowledgementConnectionSettings{
	// 							IgnoreCertificateNameMismatch: to.Ptr(true),
	// 							KeepHTTPConnectionAlive: to.Ptr(true),
	// 							SupportHTTPStatusCodeContinue: to.Ptr(true),
	// 							UnfoldHTTPHeaders: to.Ptr(true),
	// 						},
	// 						EnvelopeSettings: &armlogic.AS2EnvelopeSettings{
	// 							AutogenerateFileName: to.Ptr(true),
	// 							FileNameTemplate: to.Ptr("Test"),
	// 							MessageContentType: to.Ptr("text/plain"),
	// 							SuspendMessageOnFileNameGenerationError: to.Ptr(true),
	// 							TransmitFileNameInMimeHeader: to.Ptr(true),
	// 						},
	// 						ErrorSettings: &armlogic.AS2ErrorSettings{
	// 							ResendIfMDNNotReceived: to.Ptr(true),
	// 							SuspendDuplicateMessage: to.Ptr(true),
	// 						},
	// 						MdnSettings: &armlogic.AS2MdnSettings{
	// 							DispositionNotificationTo: to.Ptr("http://tempuri.org"),
	// 							MdnText: to.Ptr("Sample"),
	// 							MicHashingAlgorithm: to.Ptr(armlogic.HashingAlgorithmSHA1),
	// 							NeedMDN: to.Ptr(true),
	// 							ReceiptDeliveryURL: to.Ptr("http://tempuri.org"),
	// 							SendInboundMDNToMessageBox: to.Ptr(true),
	// 							SendMDNAsynchronously: to.Ptr(true),
	// 							SignMDN: to.Ptr(true),
	// 							SignOutboundMDNIfOptional: to.Ptr(true),
	// 						},
	// 						MessageConnectionSettings: &armlogic.AS2MessageConnectionSettings{
	// 							IgnoreCertificateNameMismatch: to.Ptr(true),
	// 							KeepHTTPConnectionAlive: to.Ptr(true),
	// 							SupportHTTPStatusCodeContinue: to.Ptr(true),
	// 							UnfoldHTTPHeaders: to.Ptr(true),
	// 						},
	// 						SecuritySettings: &armlogic.AS2SecuritySettings{
	// 							EnableNRRForInboundDecodedMessages: to.Ptr(true),
	// 							EnableNRRForInboundEncodedMessages: to.Ptr(true),
	// 							EnableNRRForInboundMDN: to.Ptr(true),
	// 							EnableNRRForOutboundDecodedMessages: to.Ptr(true),
	// 							EnableNRRForOutboundEncodedMessages: to.Ptr(true),
	// 							EnableNRRForOutboundMDN: to.Ptr(true),
	// 							OverrideGroupSigningCertificate: to.Ptr(false),
	// 						},
	// 						ValidationSettings: &armlogic.AS2ValidationSettings{
	// 							CheckCertificateRevocationListOnReceive: to.Ptr(true),
	// 							CheckCertificateRevocationListOnSend: to.Ptr(true),
	// 							CheckDuplicateMessage: to.Ptr(true),
	// 							CompressMessage: to.Ptr(true),
	// 							EncryptMessage: to.Ptr(false),
	// 							EncryptionAlgorithm: to.Ptr(armlogic.EncryptionAlgorithmAES128),
	// 							InterchangeDuplicatesValidityDays: to.Ptr[int32](100),
	// 							OverrideMessageProperties: to.Ptr(true),
	// 							SignMessage: to.Ptr(false),
	// 						},
	// 					},
	// 					ReceiverBusinessIdentity: &armlogic.BusinessIdentity{
	// 						Qualifier: to.Ptr("AA"),
	// 						Value: to.Ptr("AA"),
	// 					},
	// 					SenderBusinessIdentity: &armlogic.BusinessIdentity{
	// 						Qualifier: to.Ptr("ZZ"),
	// 						Value: to.Ptr("ZZ"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T22:32:54.293Z"); return t}()),
	// 		GuestIdentity: &armlogic.BusinessIdentity{
	// 			Qualifier: to.Ptr("AA"),
	// 			Value: to.Ptr("AA"),
	// 		},
	// 		GuestPartner: to.Ptr("GuestPartner"),
	// 		HostIdentity: &armlogic.BusinessIdentity{
	// 			Qualifier: to.Ptr("ZZ"),
	// 			Value: to.Ptr("ZZ"),
	// 		},
	// 		HostPartner: to.Ptr("HostPartner"),
	// 		Metadata: map[string]any{
	// 		},
	// 	},
	// }
}
