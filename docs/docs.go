// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-26 12:51:19.938646902 +0800 CST m=+0.233335062

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "YANFEI API",
        "contact": {},
        "license": {},
        "version": "0.0.1"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "测试服务器是否在线",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试"
                ],
                "summary": "PING-PONG",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/group/delete_member": {
            "get": {
                "description": "删除班组中某个成员",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "班组相关"
                ],
                "summary": "删除班组中某个成员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "班组id",
                        "name": "group_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "删除用户id",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/group/group_member": {
            "get": {
                "description": "获取班组所有成员",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "班组相关"
                ],
                "summary": "获取班组所有成员",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "班组id",
                        "name": "group_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/group/in_group": {
            "get": {
                "description": "查询自己参与的班组，包括自己创建和加入的",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "班组相关"
                ],
                "summary": "查询自己参与的班组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/group/join_group": {
            "get": {
                "description": "加入班组",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "班组相关"
                ],
                "summary": "加入班组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "群组入群口令",
                        "name": "group_key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/group/new_group": {
            "post": {
                "description": "创建新班组",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "班组相关"
                ],
                "summary": "创建新班组",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建新班组",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.GroupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/info/project_types": {
            "get": {
                "description": "获取所有工程类别",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "各种类型信息"
                ],
                "summary": "获取所有工程类别",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/info/worker_types": {
            "get": {
                "description": "获取所有工种",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "各种类型信息"
                ],
                "summary": "获取所有工种",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/record/add_hour_record": {
            "post": {
                "description": "添加工时记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作记录相关"
                ],
                "summary": "添加工时记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "工时记录数据",
                        "name": "工时记录数据",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.HourRecordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/record/add_item_record": {
            "post": {
                "description": "添加分项记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作记录相关"
                ],
                "summary": "添加分项记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "分项记录数据",
                        "name": "分项记录数据",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ItemRecordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/record/check_recorded": {
            "get": {
                "description": "检查某日是否记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作记录相关"
                ],
                "summary": "检查某日是否记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "工人id",
                        "name": "worker_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "日期",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/record/confirm_record": {
            "get": {
                "description": "确认工作记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作记录相关"
                ],
                "summary": "确认工作记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "工作记录id",
                        "name": "record_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/record/get_month_records": {
            "get": {
                "description": "查看某月的工作记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作记录相关"
                ],
                "summary": "查看某月的工作记录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "月份，形如2019-04",
                        "name": "month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/user/get_user_info": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/user/login": {
            "get": {
                "description": "小程序用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "小程序用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录码",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/user/update_user_info": {
            "post": {
                "description": "更新用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "用户个人信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.WxUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/work/publish_bao": {
            "post": {
                "description": "发布包工工作",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作相关"
                ],
                "summary": "发布包工工作",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "包工招聘数据",
                        "name": "bao_work",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.BaoWorkReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/work/publish_dian": {
            "post": {
                "description": "发布点工工作",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作相关"
                ],
                "summary": "发布点工工作",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "点工招聘数据",
                        "name": "dian_work",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.DianWorkReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/work/publish_tuji": {
            "post": {
                "description": "发布突击队工作",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作相关"
                ],
                "summary": "发布突击队工作",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "突击队招聘数据",
                        "name": "tuji_work",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.TujiWorkReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        },
        "/wx/work/search": {
            "get": {
                "description": "搜索工作，需要某个筛选参数就加上，否则可以不加；按发布时间降序排序",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "工作相关"
                ],
                "summary": "搜索工作",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "二级位置信息 选填",
                        "name": "location",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "所需工种 选填",
                        "name": "need",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "工程类别 选填",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "工作类别 选填0只返回点工和包工，1只返回突击队，默认为0",
                        "name": "work_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码，从1开始 必填",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页记录数 必填",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Message": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.BaoWorkReq": {
            "type": "object",
            "properties": {
                "construction_company": {
                    "type": "string",
                    "example": "飞燕工程队"
                },
                "desc": {
                    "type": "string",
                    "example": "包吃包住"
                },
                "final_treatment": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "location": {
                    "type": "string",
                    "example": "湖北省襄阳市"
                },
                "location_info": {
                    "type": "object",
                    "$ref": "#/definitions/model.LocationInfoReq"
                },
                "need": {
                    "type": "string",
                    "example": "钢筋工"
                },
                "project_name": {
                    "type": "string",
                    "example": "主楼建造"
                },
                "scale": {
                    "type": "string",
                    "example": "大"
                },
                "totle_price": {
                    "type": "string",
                    "example": "300"
                },
                "type": {
                    "type": "string",
                    "example": "消防"
                },
                "unit": {
                    "type": "string",
                    "example": "元/平方米"
                },
                "unit_price": {
                    "type": "string",
                    "example": "2"
                }
            }
        },
        "model.DianWorkReq": {
            "type": "object",
            "properties": {
                "construction_company": {
                    "type": "string",
                    "example": "飞燕工程队"
                },
                "desc": {
                    "type": "string",
                    "example": "包吃包住"
                },
                "final_treatment": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "location": {
                    "type": "string",
                    "example": "湖北省襄阳市"
                },
                "location_info": {
                    "type": "object",
                    "$ref": "#/definitions/model.LocationInfoReq"
                },
                "max_wage": {
                    "type": "string",
                    "example": "200"
                },
                "min_wage": {
                    "type": "string",
                    "example": "100"
                },
                "need": {
                    "type": "string",
                    "example": "钢筋工"
                },
                "project_name": {
                    "type": "string",
                    "example": "主楼建造"
                },
                "required_people": {
                    "type": "string",
                    "example": "11"
                },
                "settlement": {
                    "type": "string",
                    "example": "月薪"
                },
                "type": {
                    "type": "string",
                    "example": "消防"
                }
            }
        },
        "model.GroupRequest": {
            "type": "object",
            "properties": {
                "group_name": {
                    "type": "string"
                },
                "project_name": {
                    "type": "string"
                }
            }
        },
        "model.HourRecordRequest": {
            "type": "object",
            "properties": {
                "extra_work_hours": {
                    "type": "number",
                    "example": 1
                },
                "group_id": {
                    "type": "integer",
                    "example": 1
                },
                "record_date": {
                    "type": "string",
                    "example": "2019-05-19"
                },
                "remark": {
                    "type": "string"
                },
                "work_hours": {
                    "type": "number",
                    "example": 1.5
                },
                "worker_id": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "model.ItemRecordRequest": {
            "type": "object",
            "properties": {
                "group_id": {
                    "type": "integer",
                    "example": 1
                },
                "quantity": {
                    "type": "number",
                    "example": 1
                },
                "record_date": {
                    "type": "string",
                    "example": "2019-05-19"
                },
                "remark": {
                    "type": "string"
                },
                "subitem": {
                    "type": "string",
                    "example": "刷墙"
                },
                "unit": {
                    "type": "string",
                    "example": "平方米"
                },
                "worker_id": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "model.LocationInfoReq": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "string",
                    "example": "湖北省襄阳市樊城区武商广场店对面人民公园"
                },
                "latitude": {
                    "type": "string",
                    "example": "32.04278"
                },
                "longitude": {
                    "type": "string",
                    "example": "112.15519"
                },
                "title": {
                    "type": "string",
                    "example": "人民广场"
                }
            }
        },
        "model.TujiWorkReq": {
            "type": "object",
            "properties": {
                "construction_company": {
                    "type": "string",
                    "example": "飞燕工程队"
                },
                "desc": {
                    "type": "string",
                    "example": "包吃包住"
                },
                "final_treatment": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "location": {
                    "type": "string",
                    "example": "湖北省襄阳市"
                },
                "location_info": {
                    "type": "object",
                    "$ref": "#/definitions/model.LocationInfoReq"
                },
                "money": {
                    "type": "integer",
                    "example": 80
                },
                "need": {
                    "type": "string",
                    "example": "钢筋工"
                },
                "project_name": {
                    "type": "string",
                    "example": "主楼建造"
                },
                "required_people": {
                    "type": "integer",
                    "example": 12
                },
                "type": {
                    "type": "string",
                    "example": "消防"
                },
                "work_date": {
                    "type": "string",
                    "example": "2019-05-12"
                },
                "work_days": {
                    "type": "integer",
                    "example": 10
                },
                "work_time": {
                    "type": "string",
                    "example": "8"
                }
            }
        },
        "model.WxUserInfo": {
            "type": "object",
            "properties": {
                "hometown": {
                    "type": "string",
                    "example": "江苏"
                },
                "nick_name": {
                    "type": "string",
                    "example": "飞燕一号"
                },
                "phone": {
                    "type": "string",
                    "example": "133333"
                },
                "real_name": {
                    "type": "string",
                    "example": "张三"
                },
                "sex": {
                    "type": "string",
                    "example": "男"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
